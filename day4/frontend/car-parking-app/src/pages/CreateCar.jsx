import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const CreateCar = () => {
    const [car, setCar] = useState({
        brand: "",
        number: "",
        type: "SUV",
        incoming_time: "",
        outgoing_time: "",
        parking_slot: ""
    });
    const [alert, setAlert] = useState({ type: "", message: "" });
    const navigate = useNavigate();

    const handleChange = (e) => {
        setCar({ ...car, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post("http://localhost:8080/cars", car);
            setAlert({ type: "success", message: "Car added successfully!" });
            setTimeout(() => navigate("/"), 2000);
        } catch (error) {
            setAlert({ type: "danger", message: "Failed to add car!" });
        }
    };

    return (
        <div className="container mt-4">
            {alert.message && (
                <div className={`alert alert-${alert.type}`} role="alert">
                    {alert.message}
                </div>
            )}
            <h2>Add New Car</h2>
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label className="form-label">Car Brand</label>
                    <input type="text" name="brand" className="form-control" value={car.brand} onChange={handleChange} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Car Number</label>
                    <input type="text" name="number" className="form-control" value={car.number} onChange={handleChange} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Car Type</label>
                    <select name="type" className="form-select" value={car.type} onChange={handleChange}>
                        <option>SUV</option>
                        <option>Sedan</option>
                        <option>Hatchback</option>
                    </select>
                </div>
                <div className="mb-3">
                    <label className="form-label">Incoming Time</label>
                    <input type="time" name="incoming_time" className="form-control" value={car.incoming_time} onChange={handleChange} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Outgoing Time</label>
                    <input type="time" name="outgoing_time" className="form-control" value={car.outgoing_time} onChange={handleChange} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Parking Slot</label>
                    <input type="text" name="parking_slot" className="form-control" value={car.parking_slot} onChange={handleChange} required />
                </div>
                <button type="submit" className="btn btn-primary">Save</button>
            </form>
        </div>
    );
};

export default CreateCar;