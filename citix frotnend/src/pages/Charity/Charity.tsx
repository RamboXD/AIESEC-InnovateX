import QR from "@assets/images/download.png";
import HomeLayout from "@components/Layouts/HomeLayout";

const Charity = () => {
  return (
    <HomeLayout>
      <div className="w-full flex flex-col justify-center items-center h-full gap-4">
        <img className="w-[300px] h-[300px]" src={QR} />
        <p className="text-white font-medium text-2xl">Дом мамы «АНА ҮЙІ»</p>
        <p className="text-white font-medium text-2xl">Сделайте доброе дело!</p>
      </div>
    </HomeLayout>
  );
};

export default Charity;
