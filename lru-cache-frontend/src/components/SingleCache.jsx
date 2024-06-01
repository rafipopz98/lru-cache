import React, { useState } from "react";
import axios from "axios";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

const SingleCache = () => {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");
  const [fetcKey, setFetchKey] = useState("");
  const [duration, setDuration] = useState("");
  const [message, setMessage] = useState("");

  const calculateRemainingTime = (expiration) => {
    const expirationTime = dayjs(expiration);
    const now = dayjs();
    return expirationTime.diff(now) > 0
      ? expirationTime.fromNow(true)
      : "Expired";
  };

  const handleFetch = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/get?key=${key}`);
      setValue(response.data.value);
      setFetchKey(response.data.key);
      setDuration(response.data.expiry);
      setMessage("");
    } catch (error) {
      setValue("");
      setMessage("Error fetching cache item.");
    } finally {
      setKey();
    }
  };
  return (
    <div className="singleCache">
      <div className="title">Search in Cache</div>
      <div className="form">
        <input
          type="text"
          placeholder="Enter key here..."
          value={key}
          onChange={(e) => setKey(e.target.value)}
        />
        <button onClick={handleFetch}>
          <span className="button_top">→</span>
        </button>
      </div>
      <div className="card">
        <p className="day-text">
          <span>{value ? `Key : ${fetcKey}` : "Key Not found"}</span>
        </p>
        {value && <p className="day-text">Value : {value}</p>}
        {duration && (
          <p className="day-text">
            Duration : {calculateRemainingTime(duration)}
          </p>
        )}
      </div>
    </div>
  );
};

export default SingleCache;
