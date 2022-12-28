import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { Layout } from "../../components/Layout";
import useTitle from "../../hooks/useTitle";
import { Machine } from "../../types/machine";

interface Props {}

interface PanelProps {
  machine: Machine;
}

interface SettingsPanelProps extends PanelProps {}

const SettingsPanel: React.FC<SettingsPanelProps> = ({ machine }) => {
  useTitle(`${machine.name} - Settings`);
  return (
    <div className="p-4 bg-white rounded-lg md:p-8 dark:bg-gray-800">
      Settings!
    </div>
  );
};

interface OverviewPanelProps extends PanelProps {}

const OverviewPanel: React.FC<OverviewPanelProps> = ({ machine }) => {
  useTitle(`${machine.name} - Overview`);
  return <div className="p-4 dark:bg-gray-800">Overview!</div>;
};

interface InterfacesPanelprops extends PanelProps {}

const InterfacesPanel: React.FC<InterfacesPanelprops> = ({ machine }) => {
  useTitle(`${machine.name} - Interfaces`);

  return (
    <div className="overflow-x-auto relative shadow-md sm:rounded-lg border col-span-6 col-start-2">
      <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <th scope="col" className="py-3 px-6">
            Name
          </th>
          <th scope="col" className="py-3 px-6">
            Mac
          </th>
          <th scope="col" className="py-3 px-6">
            Addresses
          </th>
        </thead>
        <tbody>
          {machine.interfaces.map((i) => (
            <tr className="bg-white dark:bg-gray-800">
              <th
                scope="row"
                className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
              >
                {i.name}
              </th>
              <td className="py-4 px-6">{i.mac}</td>
              <td className="py-4 px-6">
                {i.addrs.map((addr) => (
                  <span className="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">
                    {addr.addr}/{addr.prefix}
                  </span>
                ))}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export const MachinePage: React.FC<Props> = () => {
  const params = useParams();
  const { data: machine, isLoading } = useQuery<Machine>(
    [`machines/${params.id}`],
    () =>
      fetch(`http://cyan.local:8090/api/machines/${params.id}/`).then((res) =>
        res.json()
      )
  );

  if (isLoading) return <div>loading...</div>;

  return (
    <Layout>
      <div className="w-full bg-white rounded-lg border shadow-md dark:bg-gray-800 dark:border-gray-700">
        <ul
          className="hidden text-sm font-medium text-center text-gray-500 rounded-lg divide-x divide-gray-200 sm:flex dark:divide-gray-600 dark:text-gray-400"
          id="fullWidthTab"
          data-tabs-toggle="#fullWidthTabContent"
          role="tablist"
        >
          <button
            id="stats-tab"
            className="text-orange-300 inline-block p-4 w-full bg-gray-50 rounted-tl-lg hover:bg-gray-100 focus:outline-none dark:bg-gray-700 dark:hover:bg-gray-600"
            type="button"
          >
            {machine!.name}
          </button>
          {/* <li className="w-full">
            <Link to={`/machines/${machine!.id}/overview`}>
              <button
                id="stats-tab"
                data-tabs-target="#stats"
                type="button"
                role="tab"
                aria-controls="stats"
                aria-selected="true"
                className={`${
                  params.page == "overview" ? "text-blue-700" : ""
                } inline-block p-4 w-full bg-gray-50 rounded-tl-lg hover:bg-gray-100 focus:outline-none dark:bg-gray-700 dark:hover:bg-gray-600`}
              >
                Overview
              </button>
            </Link>
          </li>
          <li className="w-full">
            <Link to={`/machines/${machine?.id}/settings`}>
              <button
                id="about-tab"
                data-tabs-target="#about"
                type="button"
                role="tab"
                aria-controls="about"
                aria-selected="false"
                className={`${
                  params.page == "settings" ? "text-blue-700" : ""
                } inline-block p-4 w-full bg-gray-50 hover:bg-gray-100 focus:outline-none dark:bg-gray-700 dark:hover:bg-gray-600`}
              >
                Settings
              </button>
            </Link>
          </li>
          <li className="w-full">
            <Link to={`/machines/${machine?.id}/interfaces`}>
              <button
                id="faq-tab"
                data-tabs-target="#faq"
                type="button"
                role="tab"
                aria-controls="faq"
                aria-selected="false"
                className={`${
                  params.page == "interfaces" ? "text-blue-700" : ""
                } inline-block p-4 w-full bg-gray-50 rounded-tr-lg hover:bg-gray-100 focus:outline-none dark:bg-gray-700 dark:hover:bg-gray-600`}
              >
                Interfaces
              </button>
            </Link>
          </li>
        </ul> */}
        </ul>
        <div
          id="fullWidthTabContent"
          className="border-t border-gray-200 dark:border-gray-600 p-4"
        >
          <div className="overflow-x-auto relative rounded border border-gray-700">
            <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
              <thead className="text-xs text-gray-700 uppercase dark:bg-gray-700 dark:text-gray-400">
                <th scope="col" className="py-3 px-6">
                  Name
                </th>
                <th scope="col" className="py-3 px-6">
                  Mac
                </th>
                <th scope="col" className="py-3 px-6">
                  Addresses
                </th>
              </thead>
              <tbody>
                {machine!.interfaces.map((i) => (
                  <tr className="dark:bg-gray-800">
                    <th
                      scope="row"
                      className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                    >
                      {i.name}
                    </th>
                    <td className="py-4 px-6">{i.mac}</td>
                    <td className="py-4 px-6">
                      {i.addrs.map((addr) => (
                        <span className="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">
                          {addr.addr}/{addr.prefix}
                        </span>
                      ))}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </Layout>
  );
};
