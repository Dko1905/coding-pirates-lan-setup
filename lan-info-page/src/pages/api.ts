const URL = "http://localhost:3000";

export interface Service {
  id: string,
  name: string,
  icon: string,
  darkSupported: boolean,
  online: boolean,
  hostname: string,
  port: string
}

export const getServices = async (): Promise<Service[]> => {
  return await (await fetch(URL + '/services')).json()
}
