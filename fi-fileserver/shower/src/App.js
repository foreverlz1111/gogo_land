
function App() {
  return (
      <body className="ml-4 mr-4">
      <p className="font-bold mb-2 text-xl">当前目录：.Cur if not(eq.Cur "/")
          <button onClick="pre()" className="text-md text-white bg-sky-700 hover:bg-sky-500 rounded-md">
              <i className="bi bi-arrow-bar-left"></i><span className="text-xs px-2">返回上层</span></button>
          end</p>
      <table
          className="rounded-xl border-collapse table-fixed w-full text-sm border-collapse w-full border border-slate-400 dark:border-slate-500 bg-white dark:bg-slate-800 text-sm shadow-sm hover:border-spacing-2">
          <thead className="bg-slate-50 dark:bg-slate-700">
          <tr>
              <th className=" border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left">
                  文件名称
              </th>

              <th className="w-1/4 border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left">
                  文件大小
              </th>
              <th className="w-1/4 border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left">
                  修改日期
              </th>
          </tr>
          </thead>
          <tbody>
          {/*range $Files*/}
          <tr className="border border-slate-300 dark:border-slate-700 p-4 text-slate-500 dark:text-slate-400 hover:bg-gray-200 font-md">
              <td className="text-slate-900 border">
                  <button onClick=""><span>.Name</span></button>
              </td>
              <td> 目录.大小</td>
              <td> .ModifiedTime</td>
          </tr>
          </tbody>
      </table>
      </body>
  )
}

export default App;
