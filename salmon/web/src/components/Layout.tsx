import { useNavigate } from "react-router-dom";
import { Navbar } from "./Navbar";

interface Props {
  children: React.ReactNode;
  customButton?: React.ReactNode | null;
}

export const Layout: React.FC<Props> = ({ children, customButton }) => {
  return (
    <div className="h-screen">
      <div className="dark:bg-gray-900 h-full">
        <Navbar />
        <div className="grid grid-cols-8 m-8">
          <div className="col-start-2 col-span-6">
            <div className="flex flex-row justify-end">
              {customButton != null ? <div>{customButton}</div> : null}
            </div>
            {children}
          </div>
        </div>
      </div>
    </div>
  );
};
