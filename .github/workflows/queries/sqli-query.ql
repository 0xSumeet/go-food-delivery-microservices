import go

// Find function calls to execute SQL queries where user input is concatenated into the query string.
from go.FunctionCall fc, go.Parameter p, go.BinaryExpression be
where
  (fc.getTarget().hasName("database/sql.DB.Query") or
   fc.getTarget().hasName("database/sql.DB.Exec") or
   fc.getTarget().hasName("database/sql.DB.Prepare")) and
  // Look for a binary expression where user input is concatenated into a query string.
  p.getAnArgument() = fc.getArgument(0) and
  p.getAnArgument().getAValue() = be.getLeftOperand() and
  be.getRightOperand().toString().matches(".*user_input.*")
select fc, "Potential SQL Injection risk: user input is concatenated into SQL query."

