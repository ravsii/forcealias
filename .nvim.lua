local dap = require("dap")

table.insert(dap.configurations.go, {
  type = "go",
  name = "Run {fileName}",
  request = "launch",
  program = "./cmd/forcealias/forcealias.go",
  args = function()
    -- local fName = vim.fn.input("Enter go file to run: ")
    -- local aliases = vim.fn.input("Args (leave empty for none): ")
    -- if fName == "" then
    --   vim.print("Empty input")
    --   return {}
    -- end

    -- local fName = "./pkg/analyzer/analyzer.go"
    --
    -- local aliases = "force-alias x=strings"
    --
    -- local result = { "--" }
    -- if aliases ~= "" then
    --   table.insert(result, aliases)
    -- end
    --
    -- table.insert(result, fName)

    return { "--force-alias", "x=string", "./pkg/analyzer/analyzer.go" }
  end,
  timeout = 10000,
})
