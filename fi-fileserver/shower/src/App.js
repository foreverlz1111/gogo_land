import {useRef} from "react";

function App() {

    function CurrentDir() {
        return (
            <p className="font-bold mb-2 text-xl">当前目录：.Cur if not(eq.Cur "/")
                <button className="text-md text-white bg-sky-700 hover:bg-sky-500 rounded-md">
                    <i className="bi bi-arrow-bar-left"></i><span className="text-xs px-2">返回上层</span></button>
                end
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

        return (
            <>
                <table
                    className="rounded-xl border-collapse table-fixed w-full text-sm border-collapse w-full border border-slate-400 dark:border-slate-500 bg-white dark:bg-slate-800 text-sm shadow-sm hover:border-spacing-2">
                    <thead className="bg-slate-50 dark:bg-slate-700">
                    <tr>
                        {table_menu.map((value, index) => (
                            <th className={`border border-slate-300 dark:border-slate-600 font-bold p-4 text-slate-900 dark:text-slate-200 text-left ${index < 1 ? '' : right_width}`}>{value}</th>
                        ))}
                    </tr>
                    </thead>
                    <tbody>
                    <tr className="border border-slate-300 dark:border-slate-700 p-4 text-slate-500 dark:text-slate-400 hover:bg-gray-200 font-md">
                        <td className="text-slate-900 border">
                            <button><span>我的照片.jpg</span></button>
                        </td>
                        <td className="text-slate-900 border"> 516kb</td>
                        <td className="text-slate-900 border"></td>
                    </tr>
                    </tbody>
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
