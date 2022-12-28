import { Interface } from "./interface";

export interface Machine {
  name: string;
  id: string;
  groups: string[];
  host: string;
  hostname: string;
  fqdn: string;
  interfaces: Interface[];
}
