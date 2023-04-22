export interface Invoices {
    idheader: number;
    companyname: string | null;
    address: string | null;
    numberinvoice: number | null;
    datetime: Date | null;
    createdat: Date | null;
}

export interface serverResponse  {
    count: number;
    invoices: Invoices[]
  };
  