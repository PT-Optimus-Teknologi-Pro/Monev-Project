import { useEffect, useState } from "react";
import { SwalMessage } from "../utils/SwalMessage";
import API from "../server/API";

export default function useAPISetting() {
  const [baseUrl, setBaseUrl] = useState("");
  const [apiData, setApiData] = useState<ApiConfigProps[]>([]);
  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchApi = async () => {
      try {
        const response = await API.get("/url", {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });

        setApiData(response.data.data);
        setBaseUrl(response.data.data[0].url)
      } catch (error) {
        if (error) {
          console.error("Terjadi Kesalahan");
        }
      }
    }

    fetchApi();
  }, [token]);

  const handlePostApiSetting = async () => {
    try {
      let response;
      if (apiData.length > 0) {
        response = await API.put("/url/update", {
          url: baseUrl
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
      } else {
        response = await API.post("/url/create", {
          url: baseUrl
        }, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
      }

      const message = response.data.message;
      SwalMessage({
        type: "success",
        title: "Berhasil!",
        text: message
      })

      setTimeout(() => {
        window.location.reload();
      }, 2000);
    } catch (error) {
      if (error) {
        SwalMessage({
          type: "error",
          title: "Gagal!",
          text: "Terjadi Kesalahan!"
        })
      }
    }
  }

  return {
    baseUrl,
    setBaseUrl,
    handlePostApiSetting,
    apiData
  }
}
