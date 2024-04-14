using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        var idTag = id == null ? "" : $"[{id}] - ";
        var deptTag = department == null ? "OWNER" : $"{department.ToUpper()}";

        return $"{idTag}{name} - {deptTag}";
    }
}
