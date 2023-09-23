// @ts-nocheck
import Cola from "@assets/video/cola.mp4";
import HomeLayout from "@components/Layouts/HomeLayout";
import { useEffect, useRef } from "react";
import { FaMapMarkerAlt } from "react-icons/fa";
import { useParams } from "react-router-dom";
import BounceLoader from "react-spinners/BounceLoader";
import Typewriter from "typewriter-effect";

const Portal: React.FC<MatchParams> = (props) => {
  let { roomID } = useParams();
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();

  const openCamera = async () => {
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

  useEffect(() => {
    openCamera().then((stream) => {
      userVideo.current.srcObject = stream;
      userStream.current = stream;

      webSocketRef.current = new WebSocket(
        `ws://192.168.0.149:8000/join?roomID=${roomID}`
      );

      webSocketRef.current.addEventListener("open", () => {
        webSocketRef.current.send(JSON.stringify({ join: true }));
      });

      webSocketRef.current.addEventListener("message", async (e) => {
        const message = JSON.parse(e.data);

        if (message.join) {
          callUser();
        }

        if (message.offer) {
          handleOffer(message.offer);
        }

        if (message.answer) {
          console.log("Receiving Answer");
          peerRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
        }

        if (message.iceCandidate) {
          console.log("Receiving and Adding ICE Candidate");
          try {
            await peerRef.current.addIceCandidate(message.iceCandidate);
          } catch (err) {
            console.log("Error Receiving ICE Candidate", err);
          }
        }
      });
    });
    return () => {
      // Close WebSocket connection
      if (webSocketRef.current) {
        webSocketRef.current.close();
      }

      // Stop media streams if needed
      if (userStream.current) {
        userStream.current.getTracks().forEach((track) => {
          track.stop();
        });
      }
    };
  });

  const handleOffer = async (offer) => {
    console.log("Received Offer, Creating Answer");
    peerRef.current = createPeer();

    await peerRef.current.setRemoteDescription(
      new RTCSessionDescription(offer)
    );

    userStream.current.getTracks().forEach((track) => {
      peerRef.current.addTrack(track, userStream.current);
    });

    const answer = await peerRef.current.createAnswer();
    await peerRef.current.setLocalDescription(answer);

    webSocketRef.current.send(
      JSON.stringify({ answer: peerRef.current.localDescription })
    );
  };

  const callUser = () => {
    console.log("Calling Other User");
    peerRef.current = createPeer();

    userStream.current.getTracks().forEach((track) => {
      peerRef.current.addTrack(track, userStream.current);
    });
  };

  const createPeer = () => {
    console.log("Creating Peer Connection");
    const peer = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });

    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidateEvent;
    peer.ontrack = handleTrackEvent;

    return peer;
  };

  const handleNegotiationNeeded = async () => {
    console.log("Creating Offer");

    try {
      const myOffer = await peerRef.current.createOffer();
      await peerRef.current.setLocalDescription(myOffer);

      webSocketRef.current.send(
        JSON.stringify({ offer: peerRef.current.localDescription })
      );
    } catch (err) {}
  };

  const handleIceCandidateEvent = (e) => {
    console.log("Found Ice Candidate");
    if (e.candidate) {
      console.log(e.candidate);
      webSocketRef.current.send(JSON.stringify({ iceCandidate: e.candidate }));
    }
  };

  const handleTrackEvent = (e) => {
    console.log("Received Tracks");
    partnerVideo.current.srcObject = e.streams[0];
  };

  return (
    <HomeLayout>
      <div className="w-full rounded-2xl">
        {userVideo.current && partnerVideo.current ? (
          <div className="flex flex-col gap-2">
            <video className="rounded-2xl w-full" loop autoPlay={true} muted>
              <source src={Cola} type="video/mp4" />
            </video>
            <div className="w-full text-white font-semibold text-2xl flex flex-col justify-center items-center py-4 gap-4">
              <Typewriter
                onInit={(typewriter) => {
                  typewriter
                    .typeString("Ожидание участника...")
                    .pauseFor(100)
                    .deleteAll()
                    .start();
                }}
                options={{
                  autoStart: true,
                  loop: true,
                }}
              />
              <BounceLoader color="#293447" />
            </div>
          </div>
        ) : (
          <div className="flex-col w-full relative">
            <video
              className="absolute top-0 right-3 h-1/3 w-1/3 rounded-3xl"
              autoPlay
              controls={false}
              ref={userVideo}
              muted
            ></video>
            <video
              className="rounded-2xl w-full"
              autoPlay
              controls={false}
              ref={partnerVideo}
            ></video>
            <div className="w-full mt-4 bg-info-card py-6 rounded-2xl px-5">
              <div className="flex flex-row items-center gap-2">
                <FaMapMarkerAlt size="1.5rem" color="#ffffff" />
                <p className="text-white text-xl font-semibold">
                  Ваш Собеседник из Алматы
                </p>
              </div>
              <div></div>
            </div>
          </div>
        )}
      </div>
    </HomeLayout>
  );
};

export default Portal;
