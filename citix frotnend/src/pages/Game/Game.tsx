import HomeLayout from "@components/Layouts/HomeLayout";
import axios from "axios";
import { useState } from "react";
import BounceLoader from "react-spinners/BounceLoader";

export const Game = () => {
  const [phone, setPhone] = useState<string>("");
  const [code, setCode] = useState<string>("");
  const [name, setName] = useState<string>("");
  const [tries, setTries] = useState(0);

  const [step, setStep] = useState(1);
  const [loaded, setLoaded] = useState<boolean>(true);

  const firstStep = async () => {
    await axios
      .post("http://localhost:10001/api/auth/send", { phone: phone })
      .then((res) => {
        if (res.data.status == "ok") {
          setStep(2);
        }
      });
  };

  const secondStep = async () => {
    console.log("aksndipasnidasid", phone, code);
    await axios
      .post("http://localhost:10001/api/auth/verify", {
        phone: phone,
        code: code,
      })
      .then((res) => {
        if (res.data.status === "success") {
          console.log(res.data.data.name);
          console.log("-> ", res.data.data.name);
          if (res.data.data.name === "") {
            setStep(3);
          } else if (res.data.data.name.length >= 1) {
            console.log("IM HERE");
            setStep(5);
          }
        }
      })
      .catch(() => {
        setPhone("");
      });
  };

  const thirdStep = async () => {
    await axios
      .post("http://localhost:10001/api/auth/create", {
        name: name,
        phone: phone,
      })
      .then((res) => {
        if (res.data.status === "success") {
          setStep(5);
        }
      })
      .catch(() => {
        setPhone("");
      });
  };

  return (
    <HomeLayout>
      <div className="py-1"></div>
      <div className="w-full h-full flex flex-col justify-center items-center">
        {!loaded ? (
          <BounceLoader color="#293447" size={100} />
        ) : (
          <div className="w-full h-full flex justify-center items-center mt-2">
            {step === 1 ? (
              <div className="bg-info-card rounded-2xl px-4 flex flex-col justify-center items-center w-full h-full gap-5">
                <div className="flex flex-col justify-center items-center gap-5">
                  <p className="text-white font-medium text-2xl">
                    Введите ваш номер телефона:
                  </p>
                  <input
                    value={phone}
                    onChange={(e) => setPhone(e.target.value)}
                    className="bg-info-card-active rounded-md py-3 px-3 text-white flex justify-center items-center text-xl"
                    placeholder="+7-777-777-77-77"
                    type="text"
                  />
                </div>
                <button
                  onClick={firstStep}
                  className="bg-info-card-active text-white text-2xl px-4 py-2 rounded-lg"
                >
                  Отправить
                </button>
              </div>
            ) : null}
            {step === 2 ? (
              <div className="bg-info-card rounded-2xl px-4 flex flex-col justify-center items-center w-full h-full gap-5">
                <div className="flex flex-col justify-center items-center gap-5">
                  <p className="text-white font-medium text-2xl">
                    Введите код верификации
                  </p>
                  <input
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    className="bg-info-card-active rounded-md py-3 px-3 text-white flex justify-center items-center text-xl"
                    placeholder="123456"
                    type="text"
                  />
                </div>
                <button
                  onClick={secondStep}
                  className="bg-info-card-active text-white text-2xl px-4 py-2 rounded-lg"
                >
                  Подтвердить
                </button>
              </div>
            ) : null}
            {step === 3 ? (
              <div className="bg-info-card rounded-2xl px-4 flex flex-col justify-center items-center w-full h-full gap-5">
                <div className="flex flex-col justify-center items-center gap-5">
                  <p className="text-white font-medium text-2xl">Введите имя</p>
                  <input
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    className="bg-info-card-active rounded-md py-3 px-3 text-white flex justify-center items-center text-xl"
                    placeholder="Имя"
                    type="text"
                  />
                </div>
                <button
                  onClick={thirdStep}
                  className="bg-info-card-active text-white text-2xl px-4 py-2 rounded-lg"
                >
                  Играть
                </button>
              </div>
            ) : null}
            {step == 5 ? (
              <div className="w-full">
                <div className="flex flex-col gap-3">
                  <div className=" flex flex-row w-full gap-3">
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                  </div>
                  <div className=" flex flex-row gap-3">
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                  </div>
                  <div className=" flex flex-row gap-3">
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                    <div className="py-20 w-full bg-blue-500 rounded-2xl"></div>
                  </div>
                  <div className="py-28 w-full bg-blue-500 rounded-2xl"></div>
                </div>
              </div>
            ) : null}
          </div>
        )}
      </div>
    </HomeLayout>
  );
};
