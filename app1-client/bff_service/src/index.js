const express = require('express');
const axios = require('axios');
const pkg = require('./../package.json');

// APIs versions
const enterpriseVersion = "26.00.0000_ALPHA";

require('dotenv').config();

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware para parsear el JSON recibido desde el frontend
app.use(express.json());

// ------------------------------------------------
// ------------ GETTERS FROM FRONTEND -------------
// ------------------------------------------------

// Salud del backend
app.get(
    '/health',
    (req, res) => res.json({ status: 'BFF Running', timestamp: new Date() })
);

// Comunicación con Java - Enterprise
app.get(
    `/api/v${pkg.version}/products`,
    async (req, res) => {
        try {
            const response = await axios.get(
                `http://localhost:8080/api/v${enterpriseVersion}/product`
            );

            const data = {
                brand: response.data.brand,
                model: response.data.model
            };

            res.json(data);
        } catch(e) {
            res.status(500)
                .json({ error: `Error al comunicarse con la capa de servicio: ${e}`});
        }
    }
)

app.listen(
    PORT,
    () => console.log(`BFF escuchando el puerto ${PORT}`)
);
