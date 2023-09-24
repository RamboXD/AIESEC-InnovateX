import { Home, Portal } from "@pages";
import Game from "@pages/Game/Game";
import { Games } from "@pages/Games/Games";
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
    name: "Games",
    path: "/games",
    component: <Games />,
    roles: [Role.ADMIN, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Game",
    path: "/game/:gameID",
    component: <Game />,
    roles: [Role.ADMIN, Role.ADMIN],
    isPublic: true,
  },
];
