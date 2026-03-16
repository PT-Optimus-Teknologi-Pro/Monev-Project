/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";

export default function PenyediaV5Data() {
    const [penyediaV5Data, setPenyediaV5Data] = useState<any>(null);
    const [penyediaV5Param, setPenyediaV5Param] = useState<any>(null);
    const token = localStorage.getItem('token');

    useEffect(() => {
        const fetchPenyedia = async () => {
            try {
                if (!penyediaV5Param) return;
                const response = await API.get(`/proxy/v1/penyediav5?kd_penyedia=${penyediaV5Param}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setPenyediaV5Data(response.data.data);
            } catch (error) {
                if (error) {
                    console.error("Terjadi Kesalahan");
                }
            }
        }

        fetchPenyedia();
    }, [penyediaV5Param, token]);

    return {
        penyediaV5Data,
        setPenyediaV5Param,
    }
}
