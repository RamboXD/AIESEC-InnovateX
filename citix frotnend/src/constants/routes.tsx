import { Home, Portal } from "@pages";
import Charity from "@pages/Charity/Charity";
import Dino from "@pages/Dino/Dino";
import { Game } from "@pages/Game/Game";
import { IRoute, Role } from "@ts/types";

export const ROUTES: IRoute[] = [
  {
    name: "Home",
    path: "/",
    component: <Home />,
    roles: [Role.CLIENT, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Portal",
    path: "/portal/:roomID",
    component: <Portal />,
    roles: [Role.CLIENT, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Game",
    path: "/game",
    component: <Game />,
    roles: [Role.ADMIN, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Game",
    path: "/game/dino",
    component: <Dino />,
    roles: [Role.ADMIN, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Charity",
    path: "/charity",
    component: <Charity />,
    roles: [Role.ADMIN, Role.ADMIN],
    isPublic: true,
  },
];
