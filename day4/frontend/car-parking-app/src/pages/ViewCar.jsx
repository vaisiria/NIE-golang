import React, { useState, useEffect } from "react";
import axios from "axios";
import { useParams, Link } from "react-router-dom";

const ViewCar = () => {
    const { id } = useParams();
    const [car, setCar] = useState(null);
    const [alert, setAlert] = useState({ type: "", message: "" });

    useEffect(() => {
        axios.get(`http://localhost:8080/cars/${id}`)
            .then(response => setCar(response.data.car)) // Extract 'car' from response
            .catch(() => setAlert({ type: "danger", message: "Failed to fetch car details!" }));
    }, [id]);    

    if (!car) return <p>Loading...</p>;

    return (
        <div className="container mt-4">
            {alert.message && (
                <div className={`alert alert-${alert.type}`} role="alert">
                    {alert.message}
                </div>
            )}
            <h2>Car Details</h2>
            <div className="card">
                <div className="card-body">
                    <h5 className="card-title">{car.brand} ({car.type})</h5>
                    <p><strong>Car Number:</strong> {car.number}</p>
                    <p><strong>Car Type:</strong> {car.type}</p>
                    <p><strong>Incoming Time:</strong> {car.incoming_time}</p>
                    <p><strong>Outgoing Time:</strong> {car.outgoing_time}</p>
                    <p><strong>Parking Slot:</strong> {car.parking_slot}</p>
                    <Link to="/" className="btn btn-primary">Back to List</Link>
                </div>
            </div>
        </div>
    );
};

export default ViewCar;