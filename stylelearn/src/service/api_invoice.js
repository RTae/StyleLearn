import httpClient from "@/service/httpClient";
import { server } from "@/service/constants";

export const createInvoice = async values => {
  const timeElapsed = Date.now();
  const today = new Date(timeElapsed);
  var bodyFormDataInvoice = new FormData();
  bodyFormDataInvoice.append("userID", values.userId);
  bodyFormDataInvoice.append("createDate", today.toISOString().slice(0, 10));
  bodyFormDataInvoice.append("total", values.total);
  bodyFormDataInvoice.append("detail", "");
  bodyFormDataInvoice.append("status", "0");

  const resultInvoice = await httpClient.post(
    server.INVOICE,
    bodyFormDataInvoice
  );
  if (resultInvoice.data.status === "1") {
    for (var i = 0; i < values.lesson.length; i++) {
      var bodyFormDataInvoiceLineItem = new FormData();
      bodyFormDataInvoiceLineItem.append(
        "invoiceID",
        resultInvoice.data.result
      );
      bodyFormDataInvoiceLineItem.append("lessonID", values.lesson[i].id);
      bodyFormDataInvoiceLineItem.append("quantityDay", values.lesson[i].day);
      bodyFormDataInvoiceLineItem.append(
        "amountTotal",
        parseInt(values.lesson[i].day) * 50
      );
      bodyFormDataInvoiceLineItem.append("mode", "1");
      const resultLine = await httpClient.post(
        server.INVOICE,
        bodyFormDataInvoiceLineItem
      );
      if (resultLine.data.status === "1") {
      } else {
        return resultLine.data;
      }
    }

    return resultInvoice.data;
  } else {
    return resultInvoice.data;
  }
};

export const getUnPaidInvoice = async id => {
  const result = await httpClient.get(server.INVOICE + `?userID=${id}&mode=3`);
  if (result.data.result !== null) {
    return [result.data.result, true];
  } else {
    return [result.data.result, false];
  }
};

export const cancelInvoice = async id => {
  const result = await httpClient.delete(
    server.INVOICE + `?invoiceID=${id}&mode=2`
  );
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};

export const updateStatusInvoice = async value => {
  var bodyForm = new FormData();
  bodyForm.append("invoiceID", value.invoiceID);
  bodyForm.append("status", value.status);
  bodyForm.append("mode", "2");
  const result = await httpClient.put(server.INVOICE, bodyForm);
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};

export const getLineItemInvoice = async id => {
  const result = await httpClient.get(
    server.INVOICE + `?invoiceID=${id}&mode=4`
  );
  if (result.data.status === "1") {
    return result.data;
  } else {
    return result.data;
  }
};
