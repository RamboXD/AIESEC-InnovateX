import {
  IoIosArrowDropleftCircle,
  IoIosArrowDroprightCircle,
} from "react-icons/io";
import { IoHome } from "react-icons/io5";
import { useLocation, useNavigate } from "react-router-dom";

const Footer = () => {
  const location = useLocation();
  const navigate = useNavigate();

  return (
    <div className="w-full flex flex-row justify-between items-center fixed bottom-0 bg-info-card-active py-5">
      <button
        onClick={() => {
          if (location.key !== "default") {
            navigate(-1);
          }
        }}
        disabled={location.key === "default"}
        className="w-full flex flex-row justify-center items-center"
      >
        <IoIosArrowDropleftCircle color="#ffffff" size="2.2rem" />
      </button>
      <button
        onClick={() => navigate("/")}
        className="w-full flex flex-row justify-center items-center"
      >
        <IoHome color="#ffffff" size="1.8rem" />
      </button>
      <button
        onClick={() => {
          navigate(1);
        }}
        className="w-full flex flex-row justify-center items-center"
      >
        <IoIosArrowDroprightCircle color="#ffffff" size="2.2rem" />
      </button>
    </div>
  );
};

export default Footer;
