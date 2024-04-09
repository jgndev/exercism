class Badge {
    /**
     * Formats and returns a badge string containing the provided ID, name, and department.
     * <p>
     * The badge string is formatted as follows:
     * <ul>
     *     <li>If an ID is provided, it is included in brackets at the beginning of the string.</li>
     *     <li>The name is appended next. If the name is {@code null}, this part is omitted.</li>
     *     <li>The department is appended last, in uppercase. If the department is {@code null}, "OWNER" is used as a default value.</li>
     * </ul>
     * Each segment is separated by " - ". If both ID and name are {@code null}, the string will start with the department or "OWNER".
     *
     * @param id The unique identifier for the badge, can be {@code null}.
     * @param name The name to be displayed on the badge, can be {@code null}.
     * @param department The department associated with the badge, can be {@code null}.
     * @return A formatted string representing the badge.
     */
    public String print(Integer id, String name, String department) {
        var sb = new StringBuilder();

        if (id != null) {
            sb.append("[").append(id).append("] - ");
        }

        if (name != null) {
            sb.append(name);
        }

        sb.append(" - ").append(department == null ? "OWNER" : department.toUpperCase());

        return sb.toString();
    }
}
