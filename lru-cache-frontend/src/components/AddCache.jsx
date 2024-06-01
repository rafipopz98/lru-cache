import React, { useState } from "react";
import axios from "axios";

const AddCache = ({ fetchData }) => {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");
  const [duration, setDuration] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(
        `http://localhost:8080/add?key=${key}&value=${value}&duration=${duration}`
      );
      setMessage("Cache item added successfully.");
      fetchData(); // Fetch the updated cache data
    } catch (error) {
      setMessage("Error adding cache item.");
    }
  };

  return (
    <div className="addCache">
      <form className="form" onSubmit={handleSubmit}>
        <div className="title">Add Cache</div>
        <input
          type="number"
          placeholder="key"
          value={key}
          onChange={(e) => setKey(e.target.value)}
        />
        <input
          type="text"
          placeholder="value"
          value={value}
          onChange={(e) => setValue(e.target.value)}
        />
        <input
          type="number"
          placeholder="duration in seconds"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
        />
        <button type="submit">
          <span className="button_top">Add to Cache →</span>
        </button>
      </form>
    </div>
  );
};

export default AddCache;
