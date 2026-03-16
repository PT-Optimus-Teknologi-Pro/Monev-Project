/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";
import { TahunData } from "./TahunData";

export default function KatalogV5Data() {
  const [katalogv5Data, setKatalogV5Data] = useState<KatalogV5DataProps[]>([]);
  const [katalogv5Tahun, setKatalogV5tahun] = useState<any>(TahunData[TahunData.length - 1].text);
  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchKatalogV5Data = async () => {
      try {
        const response = await API.get(`/proxy/v1/katalogv5?tahun=${katalogv5Tahun}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setKatalogV5Data(response.data.data);
      } catch (error) {
        if (error) {
          console.error("Terjadi Kesalahan");
        }
      }
    }

    fetchKatalogV5Data();
  }, [katalogv5Tahun, token]);

  return {
    katalogv5Data,
    setKatalogV5tahun,
    katalogv5Tahun
  };
}
