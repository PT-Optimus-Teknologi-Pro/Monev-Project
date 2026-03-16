/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";
import { TahunData } from "./TahunData";

export default function NonTenderData() {
  const [nonTenderData, setNonTenderData] = useState<NonTenderDataProps[]>([]);
  const [nonTenderTahun, setNonTenderTahun] = useState<any>(TahunData[TahunData.length - 1].text);
  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchNonTenderData = async () => {
      try {
        const response = await API.get(`/proxy/v1/nontenderselesai?tahun=${nonTenderTahun}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setNonTenderData(response.data.data);
      } catch (error) {
        if (error) {
          console.error("Terjadi Kesalahan");
        }
      }
    }

    fetchNonTenderData()
  }, [nonTenderTahun, token]);

  return {
    nonTenderData,
    nonTenderTahun,
    setNonTenderTahun
  }
}
