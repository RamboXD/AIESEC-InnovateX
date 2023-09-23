import HomeLayout from "@components/Layouts/HomeLayout";

export const Game = () => {
  return (
    <HomeLayout>
      <div className="flex flex-col gap-3 mt-1">
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
    </HomeLayout>
  );
};
