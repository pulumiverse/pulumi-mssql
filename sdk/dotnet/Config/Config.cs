// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Mssql
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("mssql");

        private static readonly __Value<Pulumiverse.Mssql.Config.Types.AzureAuth?> _azureAuth = new __Value<Pulumiverse.Mssql.Config.Types.AzureAuth?>(() => __config.GetObject<Pulumiverse.Mssql.Config.Types.AzureAuth>("azureAuth"));
        /// <summary>
        /// When provided, Azure AD authentication will be used when connecting.
        /// </summary>
        public static Pulumiverse.Mssql.Config.Types.AzureAuth? AzureAuth
        {
            get => _azureAuth.Get();
            set => _azureAuth.Set(value);
        }

        private static readonly __Value<string?> _hostname = new __Value<string?>(() => __config.Get("hostname") ?? Utilities.GetEnv("MSSQL_HOSTNAME"));
        /// <summary>
        /// FQDN or IP address of the SQL endpoint. Can be also set using `MSSQL_HOSTNAME` environment variable.
        /// </summary>
        public static string? Hostname
        {
            get => _hostname.Get();
            set => _hostname.Set(value);
        }

        private static readonly __Value<int?> _port = new __Value<int?>(() => __config.GetInt32("port") ?? Utilities.GetEnvInt32("MSSQL_PORT") ?? 1433);
        /// <summary>
        /// TCP port of SQL endpoint. Defaults to `1433`. Can be also set using `MSSQL_PORT` environment variable.
        /// </summary>
        public static int? Port
        {
            get => _port.Get();
            set => _port.Set(value);
        }

        private static readonly __Value<Pulumiverse.Mssql.Config.Types.SqlAuth?> _sqlAuth = new __Value<Pulumiverse.Mssql.Config.Types.SqlAuth?>(() => __config.GetObject<Pulumiverse.Mssql.Config.Types.SqlAuth>("sqlAuth"));
        /// <summary>
        /// When provided, SQL authentication will be used when connecting.
        /// </summary>
        public static Pulumiverse.Mssql.Config.Types.SqlAuth? SqlAuth
        {
            get => _sqlAuth.Get();
            set => _sqlAuth.Set(value);
        }

        public static class Types
        {

             public class AzureAuth
             {
                public string? ClientId { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? TenantId { get; set; } = null!;
            }

             public class SqlAuth
             {
                public string Password { get; set; }
                public string Username { get; set; }
            }
        }
    }
}
