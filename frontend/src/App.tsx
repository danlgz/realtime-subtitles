import { useState } from 'react';
import { Greet, StartRecord, StopRecord } from "../wailsjs/go/main/App";
import Webcam from './components/Webcam';
import { Button } from './components/Button';

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);
    const [isRecording, setIsRecording] = useState(false);

    function greet() {
        Greet(name).then(updateResultText);
    }

    const recordingToggle = () => {
      if (isRecording) {
        setIsRecording(false);
        StopRecord();
      } else {
        setIsRecording(true);
        StartRecord();
      }
    }

    return (
      <div id="App" className="flex w-full h-screen">
        <div className="bg-red-300 w-full">
          <Webcam />
        </div>
        <div className="w-full max-w-96 flex flex-col items-center py-8">
          <Button onClick={recordingToggle}>
            {isRecording ? "Stop Recording" : "Start Recording"}
          </Button>
        </div>
      </div>
    )
}

export default App
