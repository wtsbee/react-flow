import axios from "axios";
import { useQuery } from "@tanstack/react-query";

export const useQueryNodes = () => {
  const getAllNodes = async () => {
    const data = await axios({
      method: "get",
      url: `${import.meta.env.VITE_BACKEND_URL}/nodes`,
    });
    return data;
  };
  return useQuery({
    queryKey: ["nodes"],
    queryFn: getAllNodes,
    staleTime: 0,
    onError: () => {
      alert("エラー発生");
    },
  });
};
