import {useCallback, useEffect, useState} from "react";


export function GetTableBody() {
    const dir_icon = (<>&#x1F4C2; </>)
    const file_icon = (<>&#x1F4C3; </>)

    const [dir, setDir] = useState([])
    const getDir = useCallback(async () => {
        const res = await fetch("http://127.0.0.1:3001/g", {
            method: "GET",
            mode: "cors"
        })
        const resjson = await res.json()
        setDir(resjson.Files)
    }, [])
    useEffect(() => {
        let ignore = false
        if (!ignore) {
            getDir()
        }
        return () => {
            ignore = true
        }
    }, [getDir])
    return (
        <>
            <tbody>
            {dir && dir.map((v, i) => (
                <tr key={v.Name}
                    className="border border-slate-300 dark:border-slate-700 p-4 text-slate-500 dark:text-slate-400 hover:bg-gray-200 font-md">
                    <td className="text-slate-900 border">
                        <button>
                            <span>{v.IsDir ? dir_icon : file_icon}</span>
                            <span>{v.Name}</span>
                        </button>
                    </td>
                    <td className="text-slate-900 border">{v.Size}</td>
                    <td className="text-slate-900 border">{v.ModifiedTime}</td>
                </tr>
            ))
            }
            </tbody>
        </>
    )
}