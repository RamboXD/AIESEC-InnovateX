//@ts-nocheck
import Map from "@assets/images/map.jpeg";
import Video from "@assets/video/nike.mp4";
import HomeLayout from "@components/Layouts/HomeLayout";
import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function CompanyButton({ company, onClick }) {
  const gradientStyle = {
    background: `linear-gradient(70deg, ${company.primary_color}, ${company.secondary_color})`,
    animation: "gradient 10s ease infinite",
    width: "100%",
    borderRadius: "16px",
    height: "full",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  };

  return (
    <button
      style={gradientStyle}
      onClick={() => onClick(company)}
      className="background-animate bg-gradient-to-r py-12 transition-all"
    >
      <p className="text-blue-950 font-bold text-2xl">
        Выиграй приз от {company.name}
      </p>
    </button>
  );
}

function CompaniesList({ companies, onCompanyClick }) {
  const [currentCompanyIndex, setCurrentCompanyIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentCompanyIndex((prevIndex) => (prevIndex + 1) % companies.length);
    }, 3000); // 3 seconds

    return () => clearInterval(interval); // Cleanup on unmount
  }, [companies]);

  return (
    <CompanyButton
      company={companies[currentCompanyIndex]}
      onClick={() => onCompanyClick(companies[currentCompanyIndex])}
    />
  );
}

const Home: React.FC = () => {
  const [companiesData, setCompaniesData] = useState<Array<Object>>([]);
  const navigate = useNavigate();

  const create = async (e: any) => {
    e.preventDefault();

    const resp = await fetch("http://192.168.0.149:8000/createOrJoin");
    const { room_id } = await resp.json();
    navigate(`/portal/${room_id}`);
  };

  const navigateToGames = (company) => {
    navigate(`/game`, { state: { company: company } });
  };

  useEffect(() => {
    const getCompanies = async () => {
      await axios
        .get("http://localhost:10001/api/company/")
        .then((res) => {
          console.log(res);
          if (res.status === 200) {
            setCompaniesData(res.data.data);
          }
        })
        .catch(() => {});
    };

    getCompanies();
  }, []);

  return (
    <HomeLayout>
      <div className="w-full rounded-2xl bg-blue-500">
        <video className="rounded-2xl w-full" loop autoPlay={true} muted>
          <source src={Video} type="video/mp4" />
        </video>
      </div>
      <div className="w-full flex flex-col md:flex-col lg:flex-col gap-3">
        <div className="w-full flex flex-col justify-between gap-3">
          <button
            onClick={create}
            className="w-full bg-violet-700 h-full rounded-2xl py-12 flex justify-center items-center"
          >
            <p className="text-white font-bold text-3xl">PORTAL</p>
          </button>
          <div className="w-full flex flex-row md:flex-col lg:flex-col gap-3 h-full rounded-2xl">
            {companiesData.length >= 1 ? (
              <CompaniesList
                companies={companiesData}
                onCompanyClick={navigateToGames}
              />
            ) : null}
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
