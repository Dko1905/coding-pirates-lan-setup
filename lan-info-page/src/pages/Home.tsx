import { FC } from 'react';
import { Link } from 'react-router-dom';
import { nameToIcon, icons } from './icons';
import { useQuery } from "@tanstack/react-query";
import { getServices } from './api';

export const Home: FC<{}> = () => {
  // const serviceInfo = [
  //   {name: 'CS:GO 2v2', icon: 'csgo', darkSupport: true, online: true, host: 'csgo.lanparty4600.dk', port: 42},
  //   {name: 'Minecraft FFA 1', icon: 'minecraft', darkSupport: false, online: false, host: '10.42.42.xx', port: 42},
  // ];
  const { isLoading, isError, isSuccess, data: services, error } = useQuery(['services'], getServices)

  if (isLoading)
    return <span className="text-black dark:text-white">Loading...</span>
  else if (isError)
    return <span className="text-black dark:text-white">Error: {error as any}</span>

  return (
    <div className="text-black dark:text-white">
      <h1 className="text-3xl">Server info</h1>
      <div className="grid grid-cols-4 w-fit">
        {/* Header */}
        <div className="font-bold">Spil</div>
        <div className="font-bold">Name</div>
        <div className="font-bold">Status</div>
        <div className="font-bold">Host/IP</div>

        {/* Info */}
        {services.map(service => <>
          <div key={service.id + '1'} className={(service.darkSupported ? 'dark:fill-stone-300' : '') + " w-24 inline-block mr-3 self-center"}>
            {nameToIcon(service.icon)}
          </div>
          <div key={service.id + '2'} className="self-center">
            {service.name}
          </div>
          <div key={service.id + '3'} className="self-center">
            {service.online === true
            ? <span className="text-green-500 font-bold">Online</span>
            : <span className="text-red-600 font-bold">Offline</span>}
          </div>
          <div key={service.id + '4'} className="self-center">
            {service.hostname + ':' + service.port}
          </div>
        </>)}
      </div>
      <h1 className="text-3xl">Statistics</h1>
      {/* Grafana statistics using embed */}
      {/* <div className="border-4 border-red-500">
        <iframe src="https://gameserver1-grafana.lanparty4600.dk/d-solo/su8B91N4z/daniel-dashboard?orgId=1&refresh=10s&from=1667050215753&to=1667053815754&panelId=2" width="450" height="200"></iframe>
      </div> */}
    </div>
  );
};
