import { useState, useEffect } from "react";
import { createShortcode } from "../services/Shortcode.js";

export default function InputBar() {
    const [url, setUrl] = useState("");
    const [shortcode, setShortcode] = useState("");
    const [loading, setLoading] = useState(false)
    
    const handleSubmit = async (e) => {
        try {
            setLoading(true)
            const result = await createShortcode(url);
            setShortcode(result.short_code);
        } catch (error) {
            console.error(error);
        }
        setLoading(false)
    };
    
    return (
        <div>
            <input 
                type="text" 
                value={url} 
                onChange={(e) => setUrl(e.target.value)} 
                placeholder="Enter URL" 
            />
            <button type="button" onClick={handleSubmit} disabled={loading}>Shorten</button>
        {
            shortcode && (
                <div>
                    <p>Shortcode: {shortcode}</p>
                </div>
            )
        }
        </div>
    );
}