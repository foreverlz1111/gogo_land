import {useCallback, useEffect, useRef, useState} from "react";
import {GetTableBody} from "./requester";

function App() {
    let isInitial = true
    // let dir_session = ""
    let back_icon = "<"

    // function GetMyData() {
    //     const [test, setTest] = useState([])
    //     const getDir = useCallback(async () => {
    //         const res = await fetch("http://127.0.0.1:3001/g", {
    //             method: "GET",
    //             mode: "cors"
    //         })
    //         const resjson = await res.json()
    //         setTest(resjson)
    //         return resjson
    //     }, [])
    //     useEffect(() => {
    //         let ignore = false
    //         if(!ignore){
    //             getDir()
    //         }
    //         return () => {
    //             ignore = true
    //         }
    //     }, [getDir])
    //     return (
    //         <>
    //             {test.Pointer}
    //         </>
    //     )
    // }

    function CurrentDir() {
        return (
            <p className="font-bold mb-2 text-xl">当前目录：
                {isInitial ? "/" : ""}
                <button className="ml-2 text-md text-white bg-sky-700 hover:bg-sky-500 rounded-md">
                    <i className="bi bi-arrow-bar-left text-sm">
                        {back_icon}
                    </i>
                    <span className="text-xs px-2">
                        返回上层
                    </span>
                </button>
            </p>
        )
    }

    function CurrentTable() {
        const table_menu = [
            "文件名称",
            "文件大小",
            "修改日期"
        ]
        const right_width = "w-1/4"

        function TableHearder() {
            return (
                <thead className="bg-slate-50 dark:bg-slate-700">
                <tr>
                    {table_menu.map((value, index) => (
                        <th key={value + index}
                            className={`border border-slate-300 dark:border-slate-600 font-bold p-2 text-slate-900 dark:text-slate-200 text-left ${index < 1 ? '' : right_width}`}>{value}</th>
                    ))}
                </tr>
                </thead>
            )
        }

        return (
            <>
                <table
                    className="rounded-xl border-collapse table-fixed w-full text-sm border-collapse w-full border border-slate-400 dark:border-slate-500 bg-white dark:bg-slate-800 text-sm shadow-sm hover:border-spacing-2">
                    <TableHearder/>
                    <GetTableBody/>
                </table>
            </>

        )
    }

    return (
        <>

            <CurrentDir/>
            <CurrentTable/>
        </>
    )
}

export default App;
