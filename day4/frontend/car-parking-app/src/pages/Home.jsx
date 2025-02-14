import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css"; 
import "bootstrap/dist/js/bootstrap.bundle.min.js"; 
import * as bootstrap from "bootstrap"; 

const Home = () => {
    const [cars, setCars] = useState([]);
    const [alert, setAlert] = useState({ type: "", message: "" });
    const [selectedCarId, setSelectedCarId] = useState(null);

    useEffect(() => {
        fetchCars();
    }, []);

    const fetchCars = async () => {
        try {
            const response = await axios.get("http://localhost:8080/cars");
            setCars(response.data.cars);
        } catch (error) {
            setAlert({ type: "danger", message: "Failed to fetch cars!" });
        }
    };

    const confirmDelete = (id) => {
        setSelectedCarId(id);
        const modalElement = document.getElementById("deleteModal");
        if (modalElement) {
            const modal = new bootstrap.Modal(modalElement);
            modal.show();
        }
    };
    
    const deleteCar = async () => {
        if (!selectedCarId) return;
        try {
            await axios.delete(`http://localhost:8080/cars/${selectedCarId}`);
            setCars(cars.filter(car => car.id !== selectedCarId));
            setAlert({ type: "success", message: "Car deleted successfully!" });
        } catch (error) {
            setAlert({ type: "danger", message: "Failed to delete the car!" });
        } finally {
            setSelectedCarId(null);
            const modalElement = document.getElementById("deleteModal");
            if (modalElement) {
                const modal = bootstrap.Modal.getInstance(modalElement);
                if (modal) modal.hide();
            }
        }
    };    

    return (
        <div className="container mt-4">
            {alert.message && (
                <div className={`alert alert-${alert.type} alert-dismissible fade show`} role="alert">
                    {alert.message}
                    <button type="button" className="btn-close" onClick={() => setAlert({ type: "", message: "" })}></button>
                </div>
            )}

            <div className="d-flex justify-content-end mb-3">
                <Link to="/create" className="btn btn-success">Add New Car</Link>
            </div>

            <h2 className="mb-4">Car Parking List</h2>
            <table className="table table-bordered table-striped">
                <thead className="table-dark">
                    <tr>
                        <th>#</th>
                        <th>Car Brand</th>
                        <th>Car Number</th>
                        <th>Car Type</th>
                        <th>Incoming Time</th>
                        <th>Outgoing Time</th>
                        <th>Parking Slot</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {cars.map((car, index) => (
                        <tr key={car.id}>
                            <td>{index + 1}</td>
                            <td>{car.brand}</td>
                            <td>{car.number}</td>
                            <td>{car.type}</td>
                            <td>{car.incoming_time}</td>
                            <td>{car.outgoing_time}</td>
                            <td>{car.parking_slot}</td>
                            <td>
                                <Link to={`/view/${car.id}`} className="btn btn-info btn-sm">View</Link>
                                <Link to={`/edit/${car.id}`} className="btn btn-warning btn-sm mx-2">Edit</Link>
                                <button className="btn btn-danger btn-sm" onClick={() => confirmDelete(car.id)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            {/* Delete Confirmation Modal */}
            <div className="modal fade" id="deleteModal" tabIndex="-1" aria-labelledby="deleteModalLabel" aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="deleteModalLabel">Confirm Deletion</h5>
                            <button type="button" className="btn-close" data-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div className="modal-body">
                            Are you sure you want to delete this car?
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                            <button type="button" className="btn btn-danger" onClick={deleteCar}>Delete</button>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    );
};

export default Home;