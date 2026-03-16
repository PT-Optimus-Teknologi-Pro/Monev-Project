/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import API from "../server/API";
import { TahunData } from "./TahunData";

export default function TenderData() {
  const [tenderData, setTenderData] = useState<TenderDataProps[]>([]);
  const [tenderTahun, setTenderTahun] = useState<any>(TahunData[TahunData.length - 1].text);
  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchTenderData = async () => {
      try {
        const responseTenderSelesai = await API.get(`/proxy/v1/tenderselesai?tahun=${tenderTahun}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        const responseTender = await API.get(`/proxy/v1/tender?tahun=${tenderTahun}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });

        const tenderSelesaiResData = responseTenderSelesai.data.data;
        const tenderResData = responseTender.data.data;

        const tenderMap = new Map(
          tenderSelesaiResData.map((item: any) => [item.kd_tender, item])
        );

        const mergedData = tenderResData
          .filter((selesai: any) => selesai.status_tender === "Selesai")
          .map((selesai: any) => {
            const tender = tenderMap.get(selesai.kd_tender) as any;

            const {
              kd_rup_paket,
              ...restTender
            } = tender;

            return {
              ...restTender,
              ...selesai,
              kd_rup: selesai.kd_rup ?? kd_rup_paket ?? tender.kd_rup
            };
          });

        setTenderData(mergedData);
      } catch (error) {
        if (error) {
          console.error("Terjadi Kesalahan");
        }
      }
    }

    fetchTenderData()
  }, [tenderTahun, token]);

  return {
    tenderData,
    setTenderTahun,
    tenderTahun
  };
}
