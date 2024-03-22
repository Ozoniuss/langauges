

using System.Collections.Immutable;
using System.Text.Json;



using var client = new HttpClient();
try
{
    await using Stream stream = await client.GetStreamAsync("https://api.publicapis.org/entries");


    var serializerOptions = new JsonSerializerOptions();
    serializerOptions.PropertyNameCaseInsensitive = true;
    var data =
     await JsonSerializer.DeserializeAsync<JsonResponse>(stream, serializerOptions);
    if (data.Count != data.Entries.Count)
    {
        throw new Exception($"count ({data.Count}) is not equal to entries length ({data.Entries.Count})");
    }

    var currentKeys = new HashSet<string>();

    foreach (var entry in data.Entries)
    {
        var keys = entry.Keys.ToHashSet();
        if (currentKeys.Count == 0)
        {
            currentKeys = keys;
            continue;
        }
        else
        {

            if (currentKeys.Except(keys).Count() != 0 && keys.Except(currentKeys).Count() != 0)
            {
                throw new Exception("lol");
            }
        }
    }

}
catch (HttpRequestException e)
{
    Console.WriteLine("\nException Caught!");
    Console.WriteLine("Message :{0} ", e.Message);
}


class JsonResponse
{
    public int Count { get; set; }
    public List<Dictionary<string, object>> Entries { get; set; }
}