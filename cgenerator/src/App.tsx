import axios from "axios";
import React, { useEffect, useState } from "react";
import Product from "./interfaces/products";
//import logo from './logo.svg';
//import './App.css';

const App = () => {
  // State
  const [products, setProducts] = useState<Product[]>([]);
  const [_, setMyForceUpdate] = useState(0);

  // Supporting functions
  const getProductDescription = (product: Product): string => {
    let description = `${product.make} ${product.model} ${product.color} for $${product.price}`;
    product.features.forEach((feature) => {
      description += ` with ${feature}`;
    });
    description += ` and ${product.warranty} warranty`;
    return description;
  };

  const onClear = () => {
    setProducts([]);
  };

  const onLoad = async () => {
    console.info("Loading products");
    let res = await axios.get("/assets/electronics.json");
    const data = res.data;
    console.info(res.data);
    setProducts(data);
  };

  const onProcess = () => {
    console.info("Processing sales descriptions");
    products.forEach(async (product) => {
      let description = getProductDescription(product);
      let prompt = "Get a sales description for " + description;
      const postURI: string = process.env.OPENAI_ENDPOINT || "/openai";
      console.info(postURI);
      let res: any = await axios.post(
        postURI,
        {
          prompt: prompt,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      let completion = res.data.completion;
      product.description = completion;
      console.info(product);

      setProducts(products);
      setMyForceUpdate(Math.random());
    });
    console.info("Done processing sales descriptions");
  };

  // Execute a load time
  useEffect(() => {}, []);

  // Render
  return (
    <div className="App container">
      <h1>Product Description Generator</h1>
      <div>
        <button className="btn btn-primary m-2" onClick={onLoad}>
          Load
        </button>
        <button className="btn btn-secondary m-2" onClick={onProcess}>
          Process
        </button>
        <button className="btn btn-danger m-2" onClick={onClear}>
          Clear
        </button>
      </div>
      <div className="p-0">
        {products.map((product) => (
          <div className="card mb-3" key={product.id}>
            <div className="card-header bg-dark text-light fw-bold">
              {product.make} {product.model} {product.color}
            </div>
            <div className="card-body">
              <p>Description:</p>
              <p>{product.description}</p>
            </div>
            <div className="card-footer bg-light">
              Warranty: {product.warranty} Price: ${product.price.toFixed(2)}
            </div>
          </div>
        ))}
      </div>
      {/* <div>
        <table className="table table-striped">
          <thead>
            <tr>
              <th>Make</th>
              <th>Model</th>
              <th>Color</th>
              <th>Features</th>
              <th>Warranty</th>
              <th>Price</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {products.map((product) => (
              <tr key={product.id}>
                <td>{product.make}</td>
                <td>{product.model}</td>
                <td>{product.color}</td>
                <td>
                  <div
                    dangerouslySetInnerHTML={{
                      __html: product.features.join("<br/>"),
                    }}
                  />
                </td>
                <td>{product.warranty}</td>
                <td>${product.price}</td>
                <td>{product.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div> */}
    </div>
  );
};

export default App;
