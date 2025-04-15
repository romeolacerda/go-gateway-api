import { cookies } from "next/headers";

export async function getInvoice(id: string) {
    const cookiesStore = cookies();
    const apiKey = (await cookiesStore).get("apiKey")?.value;
    const response = await fetch(`http://host.docker.internal:8080/invoice/${id}`, {
      headers: {
        "X-API-KEY": apiKey as string,
      },
      cache: 'force-cache',
      next: {
        tags: [`accounts/${apiKey}/invoices/${id}`]
      }
    });
    return response.json();
  }