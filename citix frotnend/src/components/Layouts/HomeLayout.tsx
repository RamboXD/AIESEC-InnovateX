import Footer from "@components/Footer/Footer";
import Header from "@components/Header/Header";
import React from "react";

interface HomeLayoutProps {
  children: React.ReactNode;
}

const HomeLayout: React.FC<HomeLayoutProps> = ({ children }) => {
  return (
    <div className="w-full h-screen">
      <div className="w-full h-full px-4 py-4 flex flex-col gap-3">
        <Header />
        {children}
        <div className="w-full rounded-2xl py-16"></div>
      </div>
      <Footer />
    </div>
  );
};

export default HomeLayout;
