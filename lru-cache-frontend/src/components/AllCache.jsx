import React from "react";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

const AllCache = ({ data, loading, error, fetchData }) => {
  const calculateRemainingTime = (expiration) => {
    const expirationTime = dayjs(expiration);
    const now = dayjs();
    return expirationTime.diff(now) > 0
      ? expirationTime.fromNow(true)
      : "Expired";
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  console.log(data == null, "datda");

  return (
    <div className="allCache">
      <div className="top">
        <div className="title">Cache</div>
        <div
          className="refresh"
          onClick={fetchData}
        >
          &#x21bb;
        </div>
      </div>
      {data === null ? (
        <div className="card notDataFound">Not data found</div>
      ) : (
        data.map((item, index) => (
          <div className="card" key={index}>
            <p className="day-text">
              <span>Key: {item.Key}</span>
            </p>
            <p className="day-text">Value: {item.Value}</p>
            <p className="day-text">
              Expiration: {calculateRemainingTime(item.Expiration)}
            </p>
          </div>
        ))
      )}
    </div>
  );
};

export default AllCache;
