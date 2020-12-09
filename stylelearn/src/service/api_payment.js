import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const createPayment = async values => {
  const timeElapsed = Date.now();
  const today = new Date(timeElapsed);
  var bodyForm = new FormData();
  bodyForm.append("invoiceID", values.invoiceID);
  bodyForm.append("userID", values.userID);
  bodyForm.append("paymentTypeID", "a");
  bodyForm.append("status", "1");
  bodyForm.append("createDate", today.toISOString().slice(0, 10));
  bodyForm.append("tranferDate", values.tranferDate);
  bodyForm.append("amountTranfer", values.amountTranfer);
  bodyForm.append("total", values.total);
  bodyForm.append("tranferForm", values.tranferForm);
  bodyForm.append("tranferTo", values.tranferTo);

  const result = await httpClient.post(
    server.PAYMENT,
    bodyForm
  );
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};
