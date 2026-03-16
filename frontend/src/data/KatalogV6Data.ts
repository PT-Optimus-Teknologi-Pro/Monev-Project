/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";
import { TahunData } from "./TahunData";

export default function KatalogV6Data() {
  const [katalogv6Data, setKatalogV6Data] = useState<KatalogV6DataProps[]>([]);
  const [katalogv6Tahun, setKatalogV6Tahun] = useState<any>(TahunData[TahunData.length - 1].text);
  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchKatalogV6Data = async () => {
      try {
        const response = await API.get(`/proxy/v1/katalogv6?tahun=${katalogv6Tahun}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setKatalogV6Data(response.data.data);
      } catch (error) {
        if (error) {
          console.error("Terjadi Kesalahan");
        }
      }
    }

    fetchKatalogV6Data();
  }, [katalogv6Tahun, token]);

  return {
    katalogv6Data,
    setKatalogV6Tahun,
    katalogv6Tahun
  };
}
