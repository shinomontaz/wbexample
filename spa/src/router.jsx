import {createBrowserRouter} from "react-router-dom"
import Landmap from "./views/Landmap"

const router = createBrowserRouter([
    {
      path: '/',
      element: <Landmap />
    }
])

export default router;
