
export type Product = {
  ID: number;
  Name: string
  Categories: Category[];
  Price: number;
}

export type Category ={
  Id: number;
  Name: string;
}

export type Basket = {
  Id: number;
  Products: Product[];
}