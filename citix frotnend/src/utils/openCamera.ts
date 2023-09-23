export const openCamera = async () => {
  const allDevices = await navigator.mediaDevices.enumerateDevices();
  const cameras = allDevices.filter((device) => device.kind == "videoinput");
  console.log(cameras);

  const constraints = {
    audio: true,
    video: {
      deviceId: cameras[0].deviceId,
    },
  };

  try {
    return await navigator.mediaDevices.getUserMedia(constraints);
  } catch (err) {
    console.log(err);
  }
};
