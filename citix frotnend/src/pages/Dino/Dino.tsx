//@ts-nocheck
import HomeLayout from "@components/Layouts/HomeLayout";
import axios from "axios";
import React, { useEffect, useRef, useState } from "react";
import { ImExit } from "react-icons/im";
import { ColorRing } from "react-loader-spinner";
import { useLocation, useNavigate } from "react-router-dom";
import "./GameStyles.css"; // Import your CSS file here

const Dino: React.FC = () => {
  const [isJumping, setIsJumping] = useState(false);
  const [isGameOver, setIsGameOver] = useState(false);
  const [dinoPosition, setDinoPosition] = useState(150);
  const [cactusPosition, setCactusPosition] = useState(600);
  const [score, setScore] = useState<number>(1);
  const [win, setWin] = useState<boolean>(false);
  const [promo, setPromo] = useState<string>("");
  const [points, setPoints] = useState(0);

  const gameIntervalRef = useRef<NodeJS.Timeout | null>(null); // Use useRef to maintain a reference to the interval
  const [waiting, setWaiting] = useState<boolean>(false);
  const [retryLeft, setRetryLeft] = useState<boolean>(false);
  const [gameWon, setGameWon] = useState<boolean>(false);

  const navigate = useNavigate();
  console.log("now", score);

  const jump = () => {
    if (!isJumping && !isGameOver) {
      setIsJumping(true);
      setDinoPosition(200); // Increase jump height
      setTimeout(() => {
        setIsJumping(false);
        setDinoPosition(150); // Reset dino position
      }, 500);
    }
  };

  const checkCollision = () => {
    const dinoElement = document.getElementById("dino");
    const cactusElement = document.getElementById("cactus");

    if (dinoElement && cactusElement) {
      const dinoRect = dinoElement.getBoundingClientRect();
      const cactusRect = cactusElement.getBoundingClientRect();

      if (
        dinoRect.right > cactusRect.left &&
        dinoRect.left < cactusRect.right &&
        dinoRect.bottom > cactusRect.top &&
        dinoRect.top < cactusRect.bottom
      ) {
        endGame();
      }
    }
  };

  const { state } = useLocation();

  const endGame = async () => {
    console.log("--->", Math.floor(points));
    setIsGameOver(true);
    if (gameIntervalRef.current) {
      clearInterval(gameIntervalRef.current); // Clear the interval when the game is over
    }

    setWaiting(true);
    console.log({
      game_id: 1,
      user_phone: state.phone,
      company_name: state.name,
      points: score,
    });
    await axios
      .post("http://localhost:10001/api/game/play-game", {
        game_id: 1,
        user_phone: state.phone,
        company_name: state.name,
        points: score,
      })
      .then((res) => {
        console.log(res);
        if (res.data.data.game_rreport.result === "win") {
          setWin(true);
          setPromo(res.data.data.game_rreport.promocode);
        }
      });

    await axios
      .post("http://localhost:10001/api/users/is-allowed", {
        phone: state.phone,
      })
      .then((res) => {
        console.log(res);
        if (res.data.data.is_allowed === true) {
          setWaiting(false);

          setRetryLeft(true);
          setGameWon(true);
        } else {
          handleExit();
        }
      });
  };

  useEffect(() => {
    if (isGameOver && gameIntervalRef.current) {
      clearInterval(gameIntervalRef.current);
      console.log("clearing 2");
    }
  }, [isGameOver]);

  useEffect(() => {
    if (!isGameOver) {
      gameIntervalRef.current = setInterval(() => {
        checkCollision();
        setCactusPosition((prevPosition) =>
          prevPosition <= -20 ? 600 : prevPosition - 5
        );
        setScore((prevScore) => prevScore + 1 / 10);
        if (score > 0) setPoints(score);
      }, 10);

      return () => {
        if (gameIntervalRef.current) {
          clearInterval(gameIntervalRef.current);
        }
      };
    }
  }, [isGameOver]);

  useEffect(() => {
    document.addEventListener("keydown", (event) => {
      if ((event.key === " " || event.key === "ArrowUp") && !isGameOver) {
        jump();
      }
    });

    return () => {
      document.removeEventListener("keydown", (event) => {
        if ((event.key === " " || event.key === "ArrowUp") && !isGameOver) {
          jump();
        }
      });
    };
  }, [isGameOver]);

  const handleExit = () => {
    navigate("/");
  };

  const handleRestart = () => {
    setIsGameOver(false);
    setDinoPosition(150);
    setCactusPosition(600);
  };

  return (
    <HomeLayout>
      {!win ? (
        <div className="game relative w-screen">
          <div
            id="dino"
            className={isJumping ? "jump" : ""}
            style={{ bottom: `${dinoPosition}px` }}
          ></div>
          <div id="cactus" style={{ left: `${cactusPosition}px` }}></div>
          {isGameOver && (
            <div className="game-over flex justify-center items-center text-white font-VT323 relative">
              <div className="flex-col justify-center items-center">
                <h2
                  className="text-xl"
                  style={{ fontFamily: "VT323, monospace" }}
                >
                  Game Over!
                </h2>
                <div className="flex justify-center mt-2">
                  {waiting ? (
                    <ColorRing
                      visible={true}
                      height="30"
                      width="30"
                      ariaLabel="blocks-loading"
                      wrapperStyle={{}}
                      wrapperClass="blocks-wrapper"
                      colors={[
                        "#F7F7F7",
                        "#EEEEEE",
                        "#393E46",
                        "#929AAB",
                        "black",
                      ]}
                    />
                  ) : !retryLeft ? (
                    <button onClick={handleExit}>
                      <ImExit></ImExit>
                    </button>
                  ) : (
                    <button onClick={handleRestart}>
                      <svg
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M6.70406 13.5459C6.59918 13.4414 6.51597 13.3172 6.45918 13.1805C6.4024 13.0437 6.37318 12.8971 6.37318 12.7491C6.37318 12.601 6.4024 12.4544 6.45918 12.3176C6.51597 12.1809 6.59918 12.0567 6.70406 11.9522L8.95406 9.70219C9.16541 9.49084 9.45205 9.37211 9.75094 9.37211C10.0498 9.37211 10.3365 9.49084 10.5478 9.70219C10.7592 9.91353 10.8779 10.2002 10.8779 10.4991C10.8779 10.7979 10.7592 11.0846 10.5478 11.2959L10.2188 11.625H15.375V9.75C15.375 9.45163 15.4935 9.16548 15.7045 8.9545C15.9155 8.74353 16.2016 8.625 16.5 8.625C16.7984 8.625 17.0845 8.74353 17.2955 8.9545C17.5065 9.16548 17.625 9.45163 17.625 9.75V12.75C17.625 13.0484 17.5065 13.3345 17.2955 13.5455C17.0845 13.7565 16.7984 13.875 16.5 13.875H10.2188L10.5487 14.2041C10.7601 14.4154 10.8788 14.7021 10.8788 15.0009C10.8788 15.2998 10.7601 15.5865 10.5487 15.7978C10.3374 16.0092 10.0508 16.1279 9.75187 16.1279C9.45299 16.1279 9.16634 16.0092 8.955 15.7978L6.70406 13.5459ZM22.125 5.25V18.75C22.125 19.2473 21.9275 19.7242 21.5758 20.0758C21.2242 20.4275 20.7473 20.625 20.25 20.625H3.75C3.25272 20.625 2.77581 20.4275 2.42417 20.0758C2.07254 19.7242 1.875 19.2473 1.875 18.75V5.25C1.875 4.75272 2.07254 4.27581 2.42417 3.92417C2.77581 3.57254 3.25272 3.375 3.75 3.375H20.25C20.7473 3.375 21.2242 3.57254 21.5758 3.92417C21.9275 4.27581 22.125 4.75272 22.125 5.25ZM19.875 5.625H4.125V18.375H19.875V5.625Z"
                          fill="white"
                        />
                      </svg>
                    </button>
                  )}
                </div>
              </div>
            </div>
          )}

          <div
            className="absolute top-5 right-5 text-white text-lg"
            style={{ fontFamily: "VT323, monospace" }}
          >
            {" "}
            Score: {Math.floor(score)}
          </div>
        </div>
      ) : (
        <></>
      )}
    </HomeLayout>
  );
};

export default Dino;
