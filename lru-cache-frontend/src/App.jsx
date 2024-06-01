import React, { useState, useEffect } from "react";
import "./App.css";
import Header from "./components/Header";
import AddCache from "./components/AddCache";
import AllCache from "./components/AllCache";
import axios from "axios";
import SingleCache from "./components/SingleCache";

function App() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const fetchData = async () => {
    try {
      const response = await axios.get("http://localhost:8080/all");
      setData(response.data.cacheItems);
      setLoading(false);
    } catch (error) {
      setError("Error fetching cache items.");
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="container">
      <Header />
      <div className="main">
        <div className="box">
          <AddCache fetchData={fetchData} />
        </div>
        <div className="box">
          <AllCache
            data={data}
            loading={loading}
            error={error}
            fetchData={fetchData}
          />
        </div>
        <div className="box">
          <SingleCache />
        </div>
      </div>
    </div>
  );
}

export default App;
