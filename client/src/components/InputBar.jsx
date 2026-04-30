import { useState, useEffect } from "react";
import { createShortcode } from "../services/Shortcode.js";

export default function InputBar() {
  const [url, setUrl] = useState("");
  const [shortcode, setShortcode] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    try {
      setLoading(true);
      const result = await createShortcode(url);
      setShortcode(result.short_code);
    } catch (error) {
      console.error(error);
    }
    setLoading(false);
  };

  return (
    <div className="flex flex-col gap-4 w-3/4 mt-[25%]">
      <textarea
        className="resize-none w-full p-2 min-h-15 max-h-40 text-xl outline-none bg-white focus:ring-1 focus:ring-zinc-900 transition-colors border border-gray-300 rounded"
        type="text"
        value={url}
        required
        onChange={(e) => setUrl(e.target.value)}
        placeholder="Paste your link here..."
      ></textarea>
      <button
        type="button"
        onClick={handleSubmit}

        disabled={loading}
        className="bg-amber-600 p-8 h-12 text-xl font-semibold font-mono hover:shadow-lg hover:shadow-amber-30 text-white px-4 py-2 rounded"
      >
        Generate link
      </button>
      {shortcode && (
        
          <p className="text-lg mx-auto">
            <span className="font-semibold font-mono text-gray-700 text-2xl">
              Use:
            </span>{" "}
            <code className="bg-gray-100 p-2 rounded">{shortcode}</code>
          </p>

      )}
    </div>
  );
}
