import { useEffect, useRef } from "react";

const Webcam = () => {
  const videoRef = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    console.log(navigator)
    // navigator.mediaDevices.getUserMedia({
    //   audio: false,
    //   video: {
    //     frameRate: {
    //       min: 15,
    //       ideal: 30,
    //       max: 60
    //     },
    //     facingMode: 'user'
    //   }
    // }).then((stream) => {
    //   if (videoRef?.current) {
    //     videoRef.current.srcObject = stream;
    //   }
    // }).catch((err) => {
    //   alert(err)
    // })
  }, [])

  return (
    <video ref={videoRef} muted autoPlay></video>
  )
}

export default Webcam;
