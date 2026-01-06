defmodule NameBadge do
  def print(id, name, department \\ "owner") do
    fixed_department =
      if department == nil do
        "OWNER"
      else
        String.upcase(department)
      end
    if id do
      "[#{id}] - #{name} - #{fixed_department}"
    else
      "#{name} - #{fixed_department}"
    end
  end
end
