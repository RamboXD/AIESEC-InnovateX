import Map from "@assets/images/map.jpeg";
import Video from "@assets/video/nike.mp4";
import Header from "@components/Header/Header";
import HomeLayout from "@components/Layouts/HomeLayout";
import { useNavigate } from "react-router-dom";

const Home: React.FC = () => {
  const navigate = useNavigate();

  const create = async (e: any) => {
    e.preventDefault();

    const resp = await fetch("http://192.168.0.149:8000/createOrJoin");
    const { room_id } = await resp.json();
    navigate(`/portal/${room_id}`);
  };

  return (
    <HomeLayout>
      <Header />
      <div className="w-full rounded-2xl bg-blue-500">
        <video className="rounded-2xl w-full" loop autoPlay={true} muted>
          <source src={Video} type="video/mp4" />
        </video>
      </div>
      <div className="w-full flex flex-row md:flex-col lg:flex-col gap-3">
        <div className="w-full flex flex-col justify-between gap-3">
          <button
            onClick={create}
            className="w-full bg-red-500 h-full rounded-2xl md:py-24 lg:py-24 flex justify-center items-center"
          >
            <p className="text-white font-bold text-3xl">PORTAL</p>
          </button>
          <div className="w-full flex flex-row md:flex-col lg:flex-col gap-3 h-full">
            <div className="w-full bg-lime-400 rounded-2xl h-full md:py-24 lg:py-24"></div>
          </div>
        </div>
        <div className="w-full">
          <img src={Map} alt="2gis map" className="rounded-2xl" />
        </div>
      </div>
      <div className="w-full rounded-2xl bg-blue-500 py-64"></div>
    </HomeLayout>
  );
};

export default Home;
