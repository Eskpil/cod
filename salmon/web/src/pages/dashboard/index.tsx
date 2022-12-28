import { useQuery } from "@tanstack/react-query";
import { Link, useNavigate } from "react-router-dom";
import { Layout } from "../../components/Layout";
import useTitle from "../../hooks/useTitle";
import { Machine } from "../../types/machine";

interface Props {}
interface MachineElementProps {
  machine: Machine;
}

const MachineElement: React.FC<MachineElementProps> = ({ machine }) => {
  const navigate = useNavigate();

  useTitle("Machines");

  return (
    <tr
      className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50"
      onClick={() => navigate(`/machines/${machine.id}/overview`)}
    >
      <th
        scope="row"
        className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
      >
        {machine.name}
      </th>
      <td className="py-4 px-6">{machine?.host}</td>
      <td className="py-4 px-6">{machine.hostname}</td>
      <td className="py-4 px-6">
        <span className="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">
          {machine.fqdn}
        </span>
      </td>
      <td className="py-4 px-6">
        <span className="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">
          Healthy
        </span>
      </td>
    </tr>
  );
};

export const DashboardPage: React.FC<Props> = () => {
  const { data: machines, isLoading } = useQuery<Machine[]>(["machines"], () =>
    fetch("http://cyan.local:8090/api/machines/").then((res) => res.json())
  );

  if (isLoading) return <div>loading...</div>;

  return (
    <Layout
      customButton={
        <div>
          <button
            type="button"
            className="w-full text-gray-50 bg-orange-300 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none dark:focus:ring-blue-800"
          >
            New machine
          </button>
        </div>
      }
    >
      <div className="overflow-x-auto relative rounded border-x-gray-900">
        <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
          <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" className="py-3 px-6">
                Name
              </th>
              <th scope="col" className="py-3 px-6">
                Host
              </th>
              <th scope="col" className="py-3 px-6">
                Hostname
              </th>
              <th scope="col" className="py-3 px-6">
                FQDN
              </th>
              <th scope="col" className="py-3 px-6">
                Status
              </th>
            </tr>
          </thead>

          <tbody>
            {machines?.map((machine) => (
              <MachineElement machine={machine} />
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
};
