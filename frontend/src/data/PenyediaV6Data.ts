/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";

export default function PenyediaV6Data() {
    const [penyediaV6Data, setPenyediaV6Data] = useState<any>(null);
    const [penyediaV6Param, setPenyediaV6Param] = useState<any>(null);
    const token = localStorage.getItem("token");

    useEffect(() => {
        const fetchPenyedia = async () => {
            try {
                if (!penyediaV6Param) return;
                const response = await API.get(`/proxy/v1/penyediav6?kd_penyedia=${penyediaV6Param}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setPenyediaV6Data(response.data.data);
            } catch (error) {
                if (error) {
                    console.error("Terjadi Kesalahan");
                }
            }
        }

        fetchPenyedia();
    }, [penyediaV6Param, token]);

    return {
        penyediaV6Data,
        setPenyediaV6Param,
    }
}
