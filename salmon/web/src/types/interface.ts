export interface Interface {
  id: string;
  name: string;
  mac: string;
  addrs: InterfaceAddr[];
}

export interface InterfaceAddr {
  type: number;
  addr: string;
  prefix: number;
}
