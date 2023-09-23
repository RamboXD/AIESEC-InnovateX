export enum Role {
  CLIENT = "CLIENT",
  ADMIN = "ADMIN",
}

export type IRoute = {
  name: string;
  path: string;
  component: React.ReactElement;
  roles: Role[];
  isPublic: boolean;
};
