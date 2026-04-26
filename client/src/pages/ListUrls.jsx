import { useEffect, useState } from "react";
import { ListAllUrls } from "../services/Analytics";
import { analytics } from "../services/Analytics";

export default function ListUrls() {
  const [urls, setUrls] = useState([]);
  const [analyticsResponse, setAnalyticsResponse] = useState("");
  const [activeId, setActiveId] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchUrls = async () => {
      try {
        setLoading(true);
        const data = await ListAllUrls();
        setUrls(data);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    fetchUrls();
  }, []);

  const getAnalytics = async (shortcode) => {
    const parts = shortcode.split("/");
    const last = parts[parts.length - 1];
    try {
      const data = await analytics(last);
      setAnalyticsResponse(data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      {loading && (
        <div className="flex justify-center items-center h-screen">
          <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-zinc-900"></div>
        </div>
      )}
      <div className="flex justify-end gap-2">
        <div className="max-w-[80vw] mx-auto border-l-2 border-gray-200 flex flex-col items-center bg-slate-50 ">
          <div className="w-full mb-10 flex flex-row items-center gap-4 border-b border-gray-200">
            <p className="ml-3 min-w-0 w-2/3 break-words cursor-pointer text-zinc-900 text-lg font-semibold">
              Original URL
            </p>
            <h2 className="w-1/3 font-mono text-lg font-semibold">
              Short Code
            </h2>
          </div>
          {urls.map((url) => (
            <div
              className={`${
                activeId === url.short_code
                  ? "border-blue-500 bg-slate-200"
                  : "border-gray-300"
              }
            w-[80vw] p-3 flex flex-row justify-center items-center gap-4 border-b border-gray-200`}
              key={url.url_id}
            >
              <p
                className="ml-3 min-w-0 w-2/3 break-words cursor-pointer text-blue-500"
                onClick={() => window.open(url.original_url, "_blank")}
              >
                {url.original_url}
              </p>
              <h2 className="w-1/3 font-mono">{url.short_code}</h2>

              <button
                className="bg-zinc-900 text-white px-2 py-1 text-sm font-semibold rounded"
                onClick={() => {
                  getAnalytics(url.short_code);
                  setActiveId(url.short_code);
                }}
              >
                stats
              </button>
            </div>
          ))}
        </div>
        {analyticsResponse && (
          <div className="mr-1 fixed right-0 top-[30%] border border-gray-200  rounded-lg overflow-y-auto p-2 pl-4 bg-white shadow-lg">
            <h2 className="font-bold text-lg mb-2">Analytics</h2>
            {/* <p>{JSON.stringify(analyticsResponse)}</p> */}
            <p>
              <span className="font-semibold">Unique Clicks:</span>{" "}
              <span className="font-bold text-green-600">
                {analyticsResponse.unique_clicks}
              </span>
            </p>
            <p>
              <span className="font-semibold">Total Clicks:</span>{" "}
              <span className="text-green-600 font-bold">
                {analyticsResponse.total_clicks}
              </span>
            </p>
            <p>
              <span className="font-semibold">Browsers:</span>{" "}
              <span className="text-green-600 font-bold">soon</span>
            </p>
            <p>
              <span className="font-semibold">OS :</span>{" "}
              <span className="text-green-600 font-bold">soon</span>
            </p>
            <p>
              <span className="font-semibold">Referrers :</span>{" "}
              <span className="text-green-600 font-bold">soon</span>
            </p>
          </div>
        )}
      </div>
    </>
  );
}

// original_url: "http://localhost:5173/"
// short_code: "8pNEzi"
// url_id: 11
