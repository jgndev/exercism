class Badge {
    public String print(Integer id, String name, String department) {
        StringBuilder sb = new StringBuilder();

        var badgeId = id == null ? "" : String.format("[%d] - ", id);
        var badgeDept = department == null ? "OWNER" : String.format("%s", department.toUpperCase());

        sb.append(badgeId);
        sb.append(String.format("%s - ", name));
        sb.append(badgeDept);

        return sb.toString();
    }
}
