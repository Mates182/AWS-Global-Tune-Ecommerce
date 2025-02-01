using CreateInvoice.BusinessLogic;
using SoapCore;

namespace CreateInvoice
{
    public class Program
    {
        public static void Main(string[] args)
        {
            AppContext.SetSwitch("Switch.Microsoft.Data.SqlClient.DisableSqlServerPerformanceCounters", true);
            var builder = WebApplication.CreateBuilder(args);
            builder.Services.AddSoapCore();
            builder.Services.AddScoped<ISoapService, SoapService>();
            builder.Services.AddHttpContextAccessor();
            var app = builder.Build();
            app.UseRouting();
            app.UseEndpoints(endpoints =>
            {
                endpoints.UseSoapEndpoint<ISoapService>("/invoice.asmx", new SoapEncoderOptions(), SoapSerializer.XmlSerializer);
            });
            app.MapGet("/", () => "Hello World!");

            app.Run();
        }
    }
}
