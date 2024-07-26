package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
)

// record starts recording audio
func record(signal *chan int) {
	// Create a new file to store the audio recording
	fileName := fmt.Sprintf("%d.wav", time.Now().Unix())
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	// Initialize portaudio to record audio
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Open a stream to record audio
	buffer := make([]int16, 1024)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), buffer)
	if err != nil {
		fmt.Println("Error opening stream")
		return	}
	defer stream.Close()

	// Start recording audio
	err = stream.Start()
	if err != nil {
		fmt.Println("Error starting stream")
		return
	}

	// Write the WAV file header
    writeWavHeader(file, 44100, 1, 16)

    nSamples := 0
    for {
    	err = stream.Read()
    	if err != nil {
			return
     	}

      	for _, sample := range buffer {
       		err = binary.Write(file, binary.LittleEndian, sample)
       		if err != nil {
	   			return
	   		}
        }
        nSamples += len(buffer)

        select {
        case s := <-*signal:
        	fmt.Println("Signal:", s)
        	if s == 0 {
         		updateWavHeader(file, nSamples)

           		err = stream.Stop()
				if err != nil {
					fmt.Println("Error stopping stream:", err)
					return
				}

         		return
			}
		default:
        }
    }
}


func writeWavHeader(f *os.File, sampleRate, channels, bitsPerSample int) {
    f.WriteString("RIFF")
    binary.Write(f, binary.LittleEndian, uint32(0)) // Placeholder for file size
    f.WriteString("WAVE")
    f.WriteString("fmt ")
    binary.Write(f, binary.LittleEndian, uint32(16))             // Subchunk1Size for PCM
    binary.Write(f, binary.LittleEndian, uint16(1))              // AudioFormat (1 for PCM)
    binary.Write(f, binary.LittleEndian, uint16(channels))       // NumChannels
    binary.Write(f, binary.LittleEndian, uint32(sampleRate))     // SampleRate
    binary.Write(f, binary.LittleEndian, uint32(sampleRate*channels*(bitsPerSample/8))) // ByteRate
    binary.Write(f, binary.LittleEndian, uint16(channels*(bitsPerSample/8)))            // BlockAlign
    binary.Write(f, binary.LittleEndian, uint16(bitsPerSample))  // BitsPerSample
    f.WriteString("data")
    binary.Write(f, binary.LittleEndian, uint32(0)) // Placeholder for data size
}

func updateWavHeader(f *os.File, nSamples int) {
	fileSize := 44 + nSamples*2 // 44 bytes for header, 2 bytes por muestra (int16)
	dataSize := nSamples * 2
	fmt.Printf("Updating WAV header: fileSize=%d, dataSize=%d, nSamples=%d\n", fileSize, dataSize, nSamples) // Debugging info
	f.Seek(4, io.SeekStart)
	binary.Write(f, binary.LittleEndian, uint32(fileSize-8))
	f.Seek(40, io.SeekStart)
	binary.Write(f, binary.LittleEndian, uint32(dataSize))
}
